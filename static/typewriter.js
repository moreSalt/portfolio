(function () {
  var reduceMotion = window.matchMedia("(prefers-reduced-motion: reduce)").matches;
  var skipAll = false;
  var skipBtn = document.getElementById("skip");

  if (skipBtn) {
    skipBtn.addEventListener("click", function () {
      skipAll = true;
    });
  }

  function scrollToBottom() {
    window.scrollTo(0, document.documentElement.scrollHeight);
  }

  function fireTyped(root) {
    var next = root.querySelector('[hx-trigger="typed"]');
    if (next) {
      next.dispatchEvent(new Event("typed"));
    } else if (skipBtn) {
      skipBtn.style.display = "none";
    }
  }

  function dump(items) {
    items.forEach(function (item) {
      item.el.innerHTML = item.html;
      item.el.classList.remove("typing");
    });
  }

  // Human-ish cadence: long blocks type fast, short lines slower,
  // brief pauses at punctuation.
  function charDelay(text, ch) {
    if (text.length > 100) return 2 + Math.random() * 6;
    if (ch === "." || ch === "!" || ch === "?") return 120 + Math.random() * 160;
    if (ch === "," || ch === ";") return 60 + Math.random() * 80;
    return 8 + Math.random() * 24;
  }

  // Typing targets: each direct block child of the swapped-in section,
  // descending into lists so items type one by one.
  function collectTargets(root) {
    var targets = [];
    Array.prototype.slice.call(root.children).forEach(function (child) {
      if (child.hasAttribute("hx-trigger")) return;
      if (child.tagName === "UL" || child.tagName === "OL") {
        Array.prototype.slice.call(child.children).forEach(function (li) {
          targets.push(li);
        });
      } else {
        targets.push(child);
      }
    });
    return targets;
  }

  function typeAll(root) {
    if (root.tagName !== "SECTION") return;
    var targets = collectTargets(root).filter(function (el) {
      return !el.hasAttribute("data-typed");
    });
    if (!targets.length) return;

    var items = targets.map(function (el) {
      el.setAttribute("data-typed", "");
      return { el: el, html: el.innerHTML, text: el.textContent };
    });

    if (reduceMotion || skipAll) {
      dump(items);
      fireTyped(root);
      return;
    }

    items.forEach(function (item) {
      item.el.textContent = "";
    });

    var idx = 0, pos = 0;

    function step() {
      if (skipAll) {
        dump(items);
        fireTyped(root);
        return;
      }
      var cur = items[idx];
      cur.el.classList.add("typing");
      if (pos < cur.text.length) {
        var ch = cur.text[pos++];
        cur.el.textContent += ch;
        scrollToBottom();
        setTimeout(step, charDelay(cur.text, ch));
      } else {
        cur.el.innerHTML = cur.html;
        cur.el.classList.remove("typing");
        idx++;
        pos = 0;
        if (idx < items.length) {
          setTimeout(step, 250);
        } else {
          fireTyped(root);
        }
      }
    }
    setTimeout(step, 150);
  }

  document.body.addEventListener("htmx:load", function (e) {
    typeAll(e.detail.elt);
  });
})();
