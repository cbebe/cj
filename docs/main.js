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
  for (const idx in e.target.value) {
    const c = e.target.value[idx];
    if (c in data) {
      hasContent = true;
      const li = document.createElement("li");
      li.textContent = `${c} - ${data[c]} - ${toChars(data[c])} - `;
      addLinks(li, c);
      ul.appendChild(li);
    }
  }

  result.textContent = "";
  if (hasContent) {
    addLinks(result, e.target.value);
  }
});
