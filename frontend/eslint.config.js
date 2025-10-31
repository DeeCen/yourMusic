import js from '@eslint/js';
import globals from 'globals';
import vue from 'eslint-plugin-vue';
import { defineConfig } from 'eslint/config';

const conf = [
    {
        files: ['**/*.{js,mjs,cjs,vue}'],
        plugins: { js, vue },
        extends: ['js/recommended'],
        languageOptions: {
            globals: globals.browser,
            ecmaVersion: 2020,
            sourceType: 'module',
        },
        rules: {
            // 关闭组件名校验 Component name xxx should always be multi-word
            'vue/multi-word-component-names': 'off',

            // 空格
            indent: [
                'error',
                4,
                {
                    SwitchCase: 1, // switch case 缩进 1 级
                    VariableDeclarator: 1, // 变量声明缩进
                    MemberExpression: 1, // 成员表达式缩进
                    FunctionDeclaration: {
                        parameters: 1, // 函数参数缩进
                    },
                },
            ],
            quotes: [
                'error',
                'single',
                {
                    avoidEscape: true, // 允许在包含单引号的字符串中使用双引号
                    allowTemplateLiterals: true, // 允许使用模板字符串
                },
            ],
            semi: ['error', 'always'], // 必须分号
            'comma-dangle': ['error', 'never'], // 禁止尾随逗号
            'object-curly-spacing': ['error', 'always'], // 对象花括号内空格
            'array-bracket-spacing': ['error', 'never'], // 数组括号内无空格
            // 去除代码行尾的空白
            'no-trailing-spaces': 'error',

            // 花括号 { 前添加空格
            'space-before-blocks': 'error',

            // if( 括号后添加1格空格
            'keyword-spacing': ['error', { before: true, after: true }],

            'arrow-spacing': ['error', { before: true, after: true }],

            curly: ['error', 'all'], // 所有控制语句都必须使用花括号
        },
    },
    // 应用 Vue 官方推荐的规则集
    vue.configs['flat/essential'],
    vue.configs['flat/recommended'],
];
export default defineConfig(conf);
