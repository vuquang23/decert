// SPDX-License-Identifier: MIT
pragma solidity ^0.8.4;

import "./ERC721NTEnumarable.sol";
import "@openzeppelin/contracts/utils/Context.sol";

contract Decert is Context, ERC721NTEnumarable {
    address public issuer;
    address public delegatedSigner;

    uint256 private _certPtr;
    uint256 private _batchSize;

    struct Certificate {
        address issuer;
        address recipient;
        bytes32 certHash;
    }

    mapping(uint256 => Certificate) private _certData;

    mapping(bytes32 => uint256) private _hashToID;

    struct RevokedStatus {
        bool isRevoked;
        string reason;
    }

    mapping(uint256 => RevokedStatus) private _revokedStatus;

    constructor(
        address _issuer, 
        address _delegatedSigner,
        string memory _name,
        string memory _symbol    
    ) ERC721Nontransferable(_name, _symbol) {
        issuer = _issuer;
        delegatedSigner = _delegatedSigner;
    }

    modifier onlySigner() {
        require(
            _msgSender() == delegatedSigner,
            "ONLY_SIGNER_ALLOWED"
        );
        _;
    }

    function batchSize() public view returns (uint256) {
        return _batchSize;
    }

    function certData(uint256 _tokenID) public view virtual returns (Certificate memory) {
        require(!_isRevoked(_tokenID), "CERTIFICATE_REVOKED");
        return _certData[_tokenID];
    }

    function hashToID(bytes32 _certHash) public view returns (uint256) {
        return _hashToID[_certHash];
    }

    function revokedStatus(uint256 _revokedTokenID) public view virtual returns (RevokedStatus memory) {
        return _revokedStatus[_revokedTokenID];
    }

    function setBatchSize(uint256 _newBatchSize) public onlySigner {
        _batchSize = _newBatchSize;
    }

    function revokeCertificate(uint256[] memory _tokenID, string memory _reason) public virtual onlySigner {
        require(_tokenID.length <= batchSize(), "EXCEED_BATCH_SIZE");

        for (uint i = 0; i < _tokenID.length; i++) {
            require(super._exists(_tokenID[i]), "TOKEN_NOT_EXISTED");
            require(!_isRevoked(_tokenID[i]), "CERTIFICATE_REVOKED");

            _revokedStatus[_tokenID[i]] = RevokedStatus({
                isRevoked: true, 
                reason: _reason
            });

            _burn(_tokenID[i]);
        }
    }

    function batchMint(Certificate[] memory _certificate) public virtual onlySigner {
        require(_certificate.length <= batchSize(), "EXCEED_BATCH_SIZE");

        for (uint256 i = 0; i < _certificate.length; i++) {
            require(issuer == _certificate[i].issuer, "WRONG_ISSUER");
            require(!_existHash(_certificate[i].certHash), "CERT_HASH_EXISTED");

            uint256 tokenID = _increaseCertPtr();
            _mint(_certificate[i], tokenID);
        }
    }

    function _mint(Certificate memory _cert, uint256 _tokenID) internal virtual {
        super._mint(_cert.recipient, _tokenID);

        _certData[_tokenID] = _cert;
        _hashToID[_cert.certHash] = _tokenID;
    }

    function _burn(uint256 _tokenID) internal virtual {
        super._burn(_certData[_tokenID].recipient, _tokenID);
    }

    function _isRevoked(uint256 _tokenID) internal view returns (bool) {
        return _revokedStatus[_tokenID].isRevoked;
    }

    function _issuerOf(uint256 _tokenID) internal view returns (address) {
        return _certData[_tokenID].issuer;
    }

    function _existHash(bytes32 _certHash) internal view returns (bool) {
        return (_hashToID[_certHash] != 0);
    }

    function _increaseCertPtr() internal returns (uint256) {
        _certPtr += 1;
        return _certPtr;
    }
}