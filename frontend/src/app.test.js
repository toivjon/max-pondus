import App from "./app.js";
import React from 'react';
import TestRenderer from 'react-test-renderer';

test('test that Jest works', () => {
  const content = TestRenderer.create(<App />);
  const heading = content.root.findByType('h1');
  expect(heading.children[0]).toBe('Hello, World!');
});