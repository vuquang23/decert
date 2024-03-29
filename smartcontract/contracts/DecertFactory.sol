// SPDX-License-Identifier: MIT
pragma solidity ^0.8.4;

import "./Decert.sol";

contract DecertFactory {
    event DecertCreated(address indexed _scAddress);

    function CreateNewDecertCollection(
        address _issuer,
        string memory _name,
        string memory _symbol
    ) public virtual {
        Decert decert = new Decert(_issuer, _name, _symbol);
        emit DecertCreated(address(decert));
    }
}
