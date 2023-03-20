(function () {
  const body = document.body;

  const element = document.createElement('p');
  const text = document.createTextNode('Element insert from JavaScript.');
  element.appendChild(text);
  body.appendChild(element);
})();
