import React from 'react';
import TestRenderer from 'react-test-renderer';
import App from './app.js';

test('App root div has correct props', () => {
  const content = TestRenderer.create(<App />);
  const root = content.root.findByType('div');
  expect(root.props['className']).toBe("App");
});