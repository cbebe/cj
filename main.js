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
 * @type {HTMLTextAreaElement}
 */
const input = document.querySelector("#search");
const result = document.querySelector("#result");
const ul = document.querySelector("#list");

const chars = "日月金木水火土竹戈十大中一弓人心手口尸廿山女田難卜".split("");

function toChars(str) {
  return str.split("").map((e) => chars[e.codePointAt(0) - "A".codePointAt(0)])
    .join("");
}

function addLink(elem, url, name, c) {
  const a = document.createElement("a");
  a.textContent = name;
  a.href = url + encodeURIComponent(c);
  elem.appendChild(a);
}

function addLinks(elem, text) {
  addLink(elem, "plecoapi://x-callback-url/s?q=", "魚", text);
  elem.append(" - ");
  addLink(elem, "https://en.wiktionary.org/wiki/", "W", text);
}

input.addEventListener("submit", (e) => e.preventDefault());
input.addEventListener("input", (e) => {
  ul.textContent = "";
  let hasContent = false;
  for (const c in e.target.value) {
    const char = e.target.value[c];
    if (char in allChars) {
      hasContent = true;
      const li = document.createElement("li");
      li.textContent = `${char} - ${allChars[char]} - ${
        toChars(allChars[char])
      } - `;
      addLinks(li, char);
      ul.appendChild(li);
    }
  }

  result.textContent = "";
  if (hasContent) {
    addLinks(result, e.target.value);
  }
});
