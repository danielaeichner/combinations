Overview

This module creates valid swap combinations that can be used to calculate arbitrage opportunities.
In decentralized crypto there is a growing number of tokens and swap pairs available. In order to calculate arbitrage opportunities, it is necessary to first find combinations that can be traded accross exchanges.

Pairs
A pair is a smart contract that lets users swap token0 for token1 and token1 for token0. Each exchange has many pair addresses contrived of two tokens.

Pair on exchange exchange0
{
    pairAddress: pair0,
    tokens: [token0,token1]
}

Pair on exchange exchange1
{
    pairAddress: pair1,
    tokens: [token0,token1]
}


Arbitrage Opportunities
In order to find arbitrage opportunities, it is necessary to find pair combinations. These pair combinations can be an infinite amount of swaps, starting with a two swap (2swap).

Swaps
The easiest arbitrage is a 2swap:
- You own token0
- You swap token0 for token1 in pair0 (you now own token1)
- You then swap token1 for token0 in pair1
- You now own token0 again but hopefully more than before because pair0 and pair1 have different exchange rates. 

A 2swap therefore is:
{pair0 token0 -> pair0 token1 -> pair1 token1 -> pair1 token0}

With two pairs, we have a total of four different swap combinations that could be profitable
{pair0 token0 -> pair0 token1 -> pair1 token1 -> pair1 token0}
{pair0 token1 -> pair0 token0 -> pair1 token0 -> pair1 token1}
{pair1 token0 -> pair1 token1 -> pair0 token1 -> pair0 token0}
{pair1 token1 -> pair1 token0 -> pair0 token0 -> pair0 token1}

With three pairs, the 2swap combinations are:
{pair0 token0 -> pair0 token1 -> pair1 token1 -> pair1 token0}
{pair0 token1 -> pair0 token0 -> pair1 token0 -> pair1 token1}
{pair1 token0 -> pair1 token1 -> pair0 token1 -> pair0 token0}
{pair1 token1 -> pair1 token0 -> pair0 token0 -> pair0 token1}
