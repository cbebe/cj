/**
 * Downloads all the components needed for each character, grouped by length,
 * as JSON.
 *
 * Simply run in the browser console while in the Wiktionary website.
 *
 * TODO: Make it more compressed by using a custom binary format? Not really
 * sure if it's worth it considering it's not even reach a megabyte in size.
 */
const BASE_URL = "https://en.m.wiktionary.org/wiki/Appendix:Chinese_Cangjie";
const allPages = "日月金木水火土竹戈十大中一弓人心手口尸廿山女田卜難".split("")
  .map((p) => `${BASE_URL}/${encodeURIComponent(p)}`);

function getComponentMap(doc) {
  const rows = doc.querySelectorAll("table[lang=zh] tr");
  const charComponents = {};
  for (let i = 0; i < rows.length; i++) {
    const cols = Array.prototype.map.call(
      rows[i].querySelectorAll("td, th"),
      (e) => e.innerText,
    );
    const components = cols[1].substring(1, cols[1].length - 1);
    const chars = cols[2].substring(1).trim().split(", ");
    for (let j = 0; j < chars.length; j++) {
      charComponents[chars[j]] = components;
    }
  }
  return charComponents;
}

async function groupByLength() {
  const allCharComponents = {};
  await Promise.all(
    allPages.map(async (url) => {
      const html = await fetch(url).then((r) => r.text());
      const doc = new DOMParser().parseFromString(html, "text/html");
      const charComponents = getComponentMap(doc);
      for (const k in charComponents) {
        allCharComponents[k] = charComponents[k];
      }
    }),
  );
  const groups = { 1: {}, 2: {}, 3: {}, 4: {}, 5: {} };
  for (const k in allCharComponents) {
    groups[allCharComponents[k].length][k] = allCharComponents[k];
  }
  return groups;
}

function downloadObjectAsJSON(exportObj, exportName) {
  const data = encodeURIComponent(JSON.stringify(exportObj));
  const a = document.createElement("a");
  a.setAttribute("href", "data:text/json;charset=utf-8," + data);
  a.setAttribute("download", exportName + ".json");
  document.body.appendChild(a);
  a.click();
  a.remove();
}

const groups = await groupByLength();
downloadObjectAsJSON(groups, "groups");
