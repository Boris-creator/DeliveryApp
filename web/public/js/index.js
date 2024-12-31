document.addEventListener('DOMContentLoaded', () => {
    document.forms[0].addEventListener('submit', e => {
        e.preventDefault()
        const data = new FormData(e.target)
        const address = JSON.parse(data.get('address') || '{}')
        fetch('/api/v1/order', {
            method: 'POST',
            body: JSON.stringify({
                address,
                time: data.get('time'),
                comment: data.get('comment') || '',
            })
        })
    })
})