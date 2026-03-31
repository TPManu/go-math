const result = document.getElementById('result');
const form = document.querySelector('form');

form.addEventListener('submit', async (e) => {
  e.preventDefault();

  const formData = new FormData(form);

  try {
    const r = await fetch('/solve', {
      method: 'POST',
      body: formData
    });

    const text = await r.text();
    result.textContent = text;
  } catch (err) {
    console.error(err);
    result.textContent = 'Something went wrong';
  }
});
