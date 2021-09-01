// SPDX-License-Identifier: GPL-3.0

pragma solidity ^0.6.3;

abstract contract HRNG {
     function hrand() public view virtual returns (bytes32);
}

contract Demo {
    address hrngAddr;
    HRNG hrng;
    constructor (address hrngaddr) public {
        hrngAddr = hrngaddr;
        hrng = HRNG(hrngAddr);
    }

    function lastr(int p) public view returns (int) {
         bytes32 r = hrng.hrand();
         int q = p + int(r)%10;
         return q;
    }
}