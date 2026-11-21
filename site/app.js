const source = document.querySelector("#source");
const runButton = document.querySelector("#run");
const output = document.querySelector("#output");
const status = document.querySelector("#runtime-status");
const lineNumbers = document.querySelector("#line-numbers");
const wrapLines = document.querySelector("#wrap-lines");
const examples = document.querySelectorAll("[data-code]");

function updateLineNumbers() {
  lineNumbers.textContent = Array.from(
    { length: source.value.split("\n").length },
    (_, index) => index + 1,
  ).join("\n");
  lineNumbers.scrollTop = source.scrollTop;
}

async function startInterpreter() {
  if (typeof Go !== "function") {
    throw new Error("Go's WebAssembly runtime was not loaded.");
  }

  const response = await fetch("strawberry.wasm");
  if (!response.ok) {
    throw new Error(`Could not load strawberry.wasm (${response.status}).`);
  }

  const go = new Go();
  const { instance } = await WebAssembly.instantiate(
    await response.arrayBuffer(),
    go.importObject,
  );
  void go.run(instance);
  runButton.disabled = false;
  examples.forEach((example) => {
    example.disabled = false;
  });
  status.textContent = "Interpreter ready";
}

function runProgram() {
  output.classList.remove("error");

  try {
    const result = globalThis.strawberryRun(source.value);
    if (result.error) {
      output.textContent = result.error;
      output.classList.add("error");
      return;
    }
    output.textContent = result.output || "(no result)";
  } catch (error) {
    output.textContent = `Interpreter stopped: ${error.message}`;
    output.classList.add("error");
    status.textContent = "Interpreter stopped; reload to restart";
    status.classList.add("error");
    runButton.disabled = true;
  }
}

runButton.addEventListener("click", runProgram);
source.addEventListener("input", updateLineNumbers);
source.addEventListener("scroll", () => {
  lineNumbers.scrollTop = source.scrollTop;
});
wrapLines.addEventListener("change", () => {
  source.wrap = wrapLines.checked ? "soft" : "off";
});
examples.forEach((example) => {
  example.addEventListener("click", () => {
    source.value = example.dataset.code;
    examples.forEach((item) => {
      const active = item === example;
      item.classList.toggle("active", active);
      item.setAttribute("aria-pressed", active);
    });
    updateLineNumbers();
    runProgram();
  });
});
source.addEventListener("keydown", (event) => {
  if (event.key === "Enter" && (event.ctrlKey || event.metaKey)) {
    runProgram();
  }
});

updateLineNumbers();

startInterpreter().catch((error) => {
  status.textContent = error.message;
  status.classList.add("error");
  output.textContent = "Build the playground first. See the project README.";
});
