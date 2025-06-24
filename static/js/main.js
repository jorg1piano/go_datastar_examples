document.addEventListener('DOMContentLoaded', function () {
  const button = document.getElementById('click-me');
  const result = document.getElementById('result');

  if (button && result) {
    button.addEventListener('click', function () {
      result.innerHTML = '<p>Button clicked! Go-templ with air is working perfectly!</p>';
    });
  }
});
