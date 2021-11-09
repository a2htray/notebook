import typescript from 'rollup-plugin-typescript2'
import sourcemaps from 'rollup-plugin-sourcemaps'
import { nodeResolve } from '@rollup/plugin-node-resolve'
import nodePolyfills from 'rollup-plugin-polyfill-node'
import commonjs from '@rollup/plugin-commonjs'
// import { terser } from 'rollup-plugin-terser'

const rollupOptions = {
    input: './src/index.ts',
    output: {
        sourcemap: true,
        file: '../assets/js/index.js',
        format: 'iife',
        globals: {
            // 'jquery': 'jQuery',
            // 'toastr': 'toastr',
        },
        plugins: [],
    },
    external: ['gijgo'],
    plugins: [
        commonjs(),
        sourcemaps(),
        nodeResolve(),
        nodePolyfills(),
        typescript(),
    ],
    onwarn: function (warning) {
        if (warning.code === 'THIS_IS_UNDEFINED') {
            return;
        }
        console.error(warning.message);
    },
}

export default rollupOptions
