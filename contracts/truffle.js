module.exports = {
    networks: {
        development: {
            host: 'localhost',
            port: 9545,
            network_id: '*',
            gas: 8000000,
            gasPrice: 20000000000, // web3.eth.gasPrice
        },
        coverage: {
            host: 'localhost',
            port: 8555,
            network_id: '*',
            gas: 0xffffffff,
            gasPrice: 20000000000, // web3.eth.gasPrice
        },
    },
};
