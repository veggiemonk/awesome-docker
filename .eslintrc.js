module.exports = {
    env: {
        browser: true,
        node: true,
        'jest/globals': true,
    },
    extends: [
        'airbnb-base',
        'plugin:import/errors',
        'plugin:import/warnings',
        'prettier',
        'eslint:recommended',
    ],
    plugins: ['import', 'prettier', 'jest'],
    rules: {
        camelcase: 0,
        'import/order': [
            'error',
            {
                groups: ['builtin', 'external', 'parent', 'sibling', 'index'],
                'newlines-between': 'never',
            },
        ],
        'no-console': 0,
        'no-restricted-syntax': 0,
        'prefer-template': 2,
        'prettier/prettier': [
            'error',
            {
                semi: true,
                trailingComma: 'all',
                singleQuote: true,
                arrowParens: 'avoid',
                bracketSpacing: true,
                tabWidth: 4,
            },
        ],
    },
};
