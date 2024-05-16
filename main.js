let allChars = {};
async function loadJSON() {
  const data = await fetch('data/groups.json').then((r) => r.json());
  for (const i in data) {
    for (const j in data[i]) {
      allChars[j] = data[i][j];
    }
  }
}

loadJSON();

/**
 * @type {HTMLInputElement}
 */
const input = document.querySelector('input');

const ul = document.querySelector('ul');

input.addEventListener('submit', (e) => e.preventDefault());
input.addEventListener('input', (e) => {
  ul.textContent = '';
  for (const c in e.target.value) {
    const char = e.target.value[c];
    if (char in allChars) {
      const li = document.createElement('li');
      li.textContent = `${char} - ${allChars[char]} - `;
      const a = document.createElement('a');
      a.textContent = '魚';
      a.href = `plecoapi://x-callback-url/s?q=${encodeURIComponent(char)}`;
      li.appendChild(a);
      ul.appendChild(li);
    }
  }
});
