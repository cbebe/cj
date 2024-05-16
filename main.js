let allChars = {};
async function loadJSON() {
  const data = await fetch("data/groups.json").then((r) => r.json());
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
const input = document.querySelector("input");

const ul = document.querySelector("ul");

const chars = "日月金木水火土竹戈十大中一弓人心手口尸廿山女田難卜".split("");

function toChars(str) {
  return str.split("").map((e) => chars[e.codePointAt(0) - "A".codePointAt(0)])
    .join("");
}

input.addEventListener("submit", (e) => e.preventDefault());
input.addEventListener("input", (e) => {
  ul.textContent = "";
  for (const c in e.target.value) {
    const char = e.target.value[c];
    if (char in allChars) {
      const li = document.createElement("li");
      li.textContent = `${char} - ${allChars[char]} - ${
        toChars(allChars[char])
      } - `;
      {
        const a = document.createElement("a");
        a.textContent = "魚";
        a.href = `plecoapi://x-callback-url/s?q=${encodeURIComponent(char)}`;
        li.appendChild(a);
      }
      li.append(" - ");
      {
        const a = document.createElement("a");
        a.textContent = "W";
        a.href = `https://en.wiktionary.org/wiki/${encodeURIComponent(char)}`;
        li.appendChild(a);
      }

      ul.appendChild(li);
    }
  }
});
