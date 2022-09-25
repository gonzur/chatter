module.exports = {
  env: {
    browser: true,
    es2021: true,
    jest: true,
  },
  extends: [
    "plugin:react/recommended",
    "airbnb",
    "prettier",
    "plugin:storybook/recommended",
  ],
  parser: "@typescript-eslint/parser",
  parserOptions: {
    ecmaFeatures: {
      jsx: true,
    },
    ecmaVersion: "latest",
    sourceType: "module",
  },
  plugins: ["react", "@typescript-eslint", "jest"],
  rules: {
    "import/extensions": [
      "error",
      "ignorePackages",
      {
        js: "never",
        jsx: "never",
        ts: "never",
        tsx: "never",
      },
    ],
    "react/function-component-definition": [
      1,
      {
        namedComponents: "arrow-function",
      },
    ],
    "react/prop-types": [0],
    "react/jsx-filename-extension": [
      1,
      {
        extensions: [".tsx"],
      },
    ],
    "react/react-in-jsx-scope": "off",
  },
  settings: {
    "import/resolver": {
      node: {
        moduleDirectory: ["node_modules", "src"],
        extensions: [".js", ".jsx", ".ts", ".tsx"],
      },
    },
  },
};
