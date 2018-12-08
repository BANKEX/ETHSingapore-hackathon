const BankexPlasma = artifacts.require('./BankexPlasma');

module.exports = function (deployer) {
  deployer.then(async function () {
    const contract = await BankexPlasma.new();
    console.log('BankexPlasma address: ' + contract.address);
  }).catch(function () {});
};
