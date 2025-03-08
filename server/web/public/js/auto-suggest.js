function debounce(func, timeout = 300) {
    let timer
    return (...args) => {
      clearTimeout(timer)
      timer = setTimeout(() => { func.apply(this, args) }, timeout)
    }
  }
  
  const setInputValue = (input, option) => {
    const hiddenField = input.closest('.suggest__container').querySelector('.suggest__field')
    const dropdown = input.closest('.suggest__container').querySelector('.suggest__dropdown')
    const { optionValue, optionLabel } = dropdown.dataset
    let hiddenValue = null
    if (option) {
      input.value = option[optionLabel]
      input.dataset.value = option[optionLabel]
      hiddenValue = optionValue === '*' ? option : option[optionValue]
    } else {
      input.value = ''
      input.dataset.value = ''
    }
  
    if (hiddenField) {
      hiddenField.value = typeof hiddenValue === 'object' ? JSON.stringify(hiddenValue) : hiddenValue
    }
  }
  
  const hideOptions = dropdown => {
    dropdown.classList.add('suggest__dropdown--hidden')
  }
  
  const setOptions = (options, label, dropdown) => {
    dropdown.innerHTML = ''
    options.forEach(option => {
      const optionElement = document.createElement('div')
      optionElement.textContent = option[label]
      optionElement.dataset.value = JSON.stringify(option)
      optionElement.classList.add('suggest__option')
      optionElement.addEventListener('click', () => {
        const visibleField = optionElement.closest('.suggest__container').querySelector('.form-control')
        setInputValue(visibleField, option)
        hideOptions(dropdown)
      })
      dropdown.appendChild(optionElement)
    })
  }
  
  const showOptions = dropdown => {
    dropdown.classList.remove('suggest__dropdown--hidden')
    const close = e => {
      if (!e.target.closest('.suggest__container')) {
        hideOptions(dropdown)
        document.removeEventListener('click', close)
      }
    }
    document.addEventListener('click', close)
  }
  
  const setInputHandlers = input => {
    if (input.dataset.inited) return
    input.addEventListener('focus', () => {
      showOptions(input.closest('.suggest__container').querySelector('.suggest__dropdown'))
    })
    input.addEventListener('blur', () => {
      if (input.value.length) {
        input.value = input.dataset.value || ''
      } else {
        setInputValue(input, null)
      }
    })
    input.dataset.inited = '1'
  }
  
  document.addEventListener('input', debounce(e => {
    const { url } = e.target.dataset
    if (!url) return
  
    const dropdown = e.target.closest('.suggest__container').querySelector('.suggest__dropdown')
  
    const query = e.target.value.trim()
    if (!query.length) {
      setOptions([], '', dropdown)
      return
    }
  
    setInputHandlers(e.target)
  
    const requestSchema = {
        request: {
            query: "",
            highestToponym: "region",
            lowestToponym: "house",
            searchInsideAddress: null,
            language: "ru",
            priorityAddress: null,
            countries: [
                "RU"
            ]
        },
        query: "query",
    }

    requestSchema.request[requestSchema.query] = query
  
    const getFieldValueByName = fieldName => document.querySelector(`[name="${fieldName}"]`).value
    const requiredRule = (val, isMultiple) => (isMultiple ? val.length && val.every(value => !!value) : !!val)
  
    Object.entries(requestSchema.values || {}).forEach(([key, { field, required }]) => {
      const isMultiple = Array.isArray(field)
      const val = isMultiple ? field.map(getFieldValueByName) : getFieldValueByName(field)
  
      if (required && !requiredRule(val, isMultiple)) {
        setOptions([], '', dropdown)
        return
      }
  
      if (field) {
        requestSchema.request[key] = val
      }
    })
  
    fetch(url, {
      method: 'POST',
      body: JSON.stringify(requestSchema.request),
      headers: {
        'Content-Type': 'application/json',
        Accept: 'application/json',
      },
    }).then(response => response.json())
      .then(data => {
        setOptions(data, dropdown.dataset.optionLabel, dropdown)
        showOptions(dropdown)
      })
  }))
  