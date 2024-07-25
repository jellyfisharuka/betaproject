document.getElementById('generate-form').addEventListener('submit', async function(event) {
    event.preventDefault();

    const prompt = document.getElementById('prompt').value;
    const responseContainer = document.getElementById('response');

    const response = await fetch('/api/generate', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/x-www-form-urlencoded',
        },
        body: new URLSearchParams({
            'prompt': prompt
        })
    });

    if (response.ok) {
        const text = await response.text();
        responseContainer.innerText = text;
    } else {
        responseContainer.innerText = 'Error: Unable to generate content';
    }
});
