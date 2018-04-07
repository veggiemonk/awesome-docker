module.exports = {
    "env": {
        "browser": true,
        "node": true
    },
    "extends": [
        "airbnb-base",
        "plugin:import/errors",
        "plugin:import/warnings",
        "prettier",
      ],
      "plugins": [
        "import",
        "prettier",
      ],
      "rules": {
        "import/order": ["error", {
          "groups": ["builtin", "external", "parent", "sibling", "index"],
          "newlines-between": "never"
          }],
        "no-console": 0,
        "prefer-template": 2,
        "prettier/prettier": [
          "error",
          {
            "singleQuote": true,
            "trailingComma": "all",
          }
        ]
      }
};