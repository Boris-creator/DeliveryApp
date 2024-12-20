console.log('It\'s here beacuse "index.js" is included in IndexNames');

fetch('http://localhost:3000/api/v1/suggest',
{
    method: 'POST',
    headers: {
        'Content-Type': 'application/json',
    },
    body: JSON.stringify({
        query: 'test'
    })
}
)