module.exports = {
    devServer: {
        port: 3000,
        proxy: {
            '/api': {
                target: 'http://localhost:8080', // 后端服务器地址
                changeOrigin: true, pathRewrite: {
                    '^/api': '' 
                }
            }
        }
    },
    configureWebpack: {
        resolve: {
            alias: {
                '@': require('path').resolve(__dirname, 'src')
            }
        },
    }
}

