document.addEventListener('DOMContentLoaded', () => {
    document.forms[0].addEventListener('submit', e => {
        e.preventDefault()
        const data = new FormData(e.target)
        const address = JSON.parse(data.get('address') || '{}')
        data.delete('address')
        data.set('address[fullAddress]', address.fullAddress)
        data.set('address[geo_lat]', address.data.geo_lat)
        data.set('address[geo_lon]', address.data.geo_lon)
        fetch('/api/v1/order', {
            method: 'POST',
            body: data,
        })
    })
})