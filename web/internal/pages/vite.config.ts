import { defineConfig, loadEnv } from 'vite'
import path from 'path'
import vue from '@vitejs/plugin-vue'
import { quasar, transformAssetUrls } from '@quasar/vite-plugin'

const ASSET_URL = process.env.ASSET_URL || '/olympus';

// https://vitejs.dev/config/
export default ({ mode }) => {
    process.env = {...process.env, ...loadEnv(mode, process.cwd())};

    return defineConfig({
        base: `/`,

        define: {
            'process.env': process.env
        },

    	plugins: [
    		vue({
    			template: { transformAssetUrls }
    		}),
    		quasar({
    			scssVariables: 'src/styles/quasar.scss',
    			sassVariables: 'src/styles/quasar.sass'
    		})
    	],

        resolve: {
            alias: {
                '@': path.resolve(__dirname, './src'),
                'src': path.resolve(__dirname, './src'),
            },
        },

        server: {
            port: "3000",
            host: "0.0.0.0",
        }
    });
}
