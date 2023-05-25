package main

import (
	"context"
	"fmt"
	"forgeLiquidation/discord"
	"forgeLiquidation/erc20"
	"forgeLiquidation/flashloan"
	"forgeLiquidation/forge"
	"forgeLiquidation/liquidationFinder"
	"forgeLiquidation/multicall"
	"forgeLiquidation/unirouterv2"
	"log"
	"math/big"
	"os"
	"strings"
	"time"

	"github.com/gofor-little/env"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func findXauLiquidation(liquidationFinderContract *liquidationFinder.LiquidationFinder, skip *big.Int) liquidationFinder.LiquidationFinderLiquidation {
	liquidationOportunities, err := liquidationFinderContract.GetLiquidation(&bind.CallOpts{}, big.NewInt(0), big.NewInt(119000000), skip)
	if err != nil {
		log.Fatal(err)
	}
	return liquidationOportunities

}

func findWtiLiquidation(liquidationFinderContract *liquidationFinder.LiquidationFinder, skip *big.Int) liquidationFinder.LiquidationFinderLiquidation {
	liquidationOportunities, err := liquidationFinderContract.GetLiquidation(&bind.CallOpts{}, big.NewInt(1), big.NewInt(124000000), skip)
	if err != nil {
		log.Fatal(err)
	}
	return liquidationOportunities

}

func findCnyLiquidation(liquidationFinderContract *liquidationFinder.LiquidationFinder, skip *big.Int) liquidationFinder.LiquidationFinderLiquidation {
	liquidationOportunities, err := liquidationFinderContract.GetLiquidation(&bind.CallOpts{}, big.NewInt(2), big.NewInt(109000000), skip)
	if err != nil {
		log.Fatal(err)
	}
	return liquidationOportunities

}

func findXagLiquidation(liquidationFinderContract *liquidationFinder.LiquidationFinder, skip *big.Int) liquidationFinder.LiquidationFinderLiquidation {
	liquidationOportunities, err := liquidationFinderContract.GetLiquidation(&bind.CallOpts{}, big.NewInt(3), big.NewInt(124000000), skip)
	if err != nil {
		log.Fatal(err)
	}
	return liquidationOportunities

}

func findEurLiquidation(liquidationFinderContract *liquidationFinder.LiquidationFinder, skip *big.Int) liquidationFinder.LiquidationFinderLiquidation {
	liquidationOportunities, err := liquidationFinderContract.GetLiquidation(&bind.CallOpts{}, big.NewInt(4), big.NewInt(109000000), skip)
	if err != nil {
		log.Fatal(err)
	}
	return liquidationOportunities

}

func findJpyLiquidation(liquidationFinderContract *liquidationFinder.LiquidationFinder, skip *big.Int) liquidationFinder.LiquidationFinderLiquidation {
	liquidationOportunities, err := liquidationFinderContract.GetLiquidation(&bind.CallOpts{}, big.NewInt(5), big.NewInt(109500000), skip)
	if err != nil {
		log.Fatal(err)
	}
	return liquidationOportunities

}

func findTslaLiquidation(liquidationFinderContract *liquidationFinder.LiquidationFinder, skip *big.Int) liquidationFinder.LiquidationFinderLiquidation {
	liquidationOportunities, err := liquidationFinderContract.GetLiquidation(&bind.CallOpts{}, big.NewInt(6), big.NewInt(144900000), skip)
	if err != nil {
		log.Fatal(err)
	}
	return liquidationOportunities
}

func addApprove(tos *[]common.Address, data *[][]byte, amount *big.Int, token common.Address, addressToApprove common.Address) {
	erc20parsed, err := abi.JSON(strings.NewReader(erc20.Erc20ABI))
	if err != nil {
		log.Fatal(err)
	}
	approveBytes, err := erc20parsed.Pack("approve", addressToApprove, amount)
	if err != nil {
		log.Fatal(err)
	}
	*tos = append(*tos, token)
	*data = append(*data, approveBytes)
}

func addTransfer(tos *[]common.Address, data *[][]byte, amount *big.Int, token common.Address, addressToTransfer common.Address) {
	erc20parsed, err := abi.JSON(strings.NewReader(erc20.Erc20ABI))
	if err != nil {
		log.Fatal(err)
	}
	approveBytes, err := erc20parsed.Pack("transfer", addressToTransfer, amount)
	if err != nil {
		log.Fatal(err)
	}
	*tos = append(*tos, token)
	*data = append(*data, approveBytes)
}

func addSwap(tos *[]common.Address, data *[][]byte, amountIn *big.Int, path []common.Address, recipient common.Address, deadline *big.Int, routerAdrress common.Address, routerInstance *unirouterv2.Uniswapv2router) *big.Int {
	swapParsed, err := abi.JSON(strings.NewReader(unirouterv2.Uniswapv2routerABI))
	if err != nil {
		fmt.Println("error parsing abi swap")
		log.Fatal(err)
	}

	amountsOut, err := routerInstance.GetAmountsOut(&bind.CallOpts{}, amountIn, path)
	if err != nil {
		fmt.Println("error getting amounts out")
		log.Fatal(err)
	}
	minammount := big.NewInt(0)
	swapBytes, err := swapParsed.Pack("swapExactTokensForTokens", amountIn, minammount, path, recipient, deadline)
	if err != nil {
		fmt.Println("error packing swap")
		log.Fatal(err)
	}

	*tos = append(*tos, routerAdrress)
	*data = append(*data, swapBytes)
	return amountsOut[1]
}

func addLiquidation(tos *[]common.Address, data *[][]byte, accountToLiquidate common.Address, assetId *big.Int, ammountToLiquidate *big.Int, forgeAdress common.Address, forgeInstance *forge.Forge) {
	forgeParsed, err := abi.JSON(strings.NewReader(forge.ForgeABI))
	if err != nil {
		fmt.Println("error parsing abi forge")
		log.Fatal(err)
	}

	liquidationBytes, err := forgeParsed.Pack("liquidate", accountToLiquidate, assetId, ammountToLiquidate)
	if err != nil {
		fmt.Println("error packing liquidation")
		log.Fatal(err)
	}

	*tos = append(*tos, forgeAdress)
	*data = append(*data, liquidationBytes)
}

func flashloanAndLiquidate(ammountToLiquidate *big.Int, conn *ethclient.Client, collateralAdress common.Address, tokenAddress common.Address, contractAdress common.Address, assetID *big.Int, accountToLiquidate common.Address, syntethicName string, skip *big.Int, privateKeyString string, publicKeyString string, discordUrl string) {
	start := time.Now()

	fmt.Println("starting flashloan and liquidation for ", syntethicName)
	fmt.Println("ammount to liquidate", ammountToLiquidate)
	fmt.Println("account to liquidate", accountToLiquidate)
	fmt.Println("collateral adress", collateralAdress)
	fmt.Println("token adress", tokenAddress)
	fmt.Println("contract adress", contractAdress)
	fmt.Println("asset id", assetID)

	forgeAddress := common.HexToAddress("0x4938D2016e7446a24b07635611bD34289Df42ECb")
	forgeContract, err := forge.NewForge(forgeAddress, conn)
	if err != nil {
		log.Fatal(err)
	}

	var path []common.Address
	path = append(path, collateralAdress)
	path = append(path, tokenAddress)

	forgeRouterAddress := common.HexToAddress("0x1ca3ec7C382183138134B9da66D36F6c042d8DAD")
	forgeRouterContract, err := unirouterv2.NewUniswapv2router(forgeRouterAddress, conn)
	if err != nil {
		log.Fatal(err)
	}

	ammountIn, err := forgeRouterContract.GetAmountsIn(&bind.CallOpts{}, ammountToLiquidate, path)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("ammount in", ammountIn[0])

	var data [][]byte
	var tos []common.Address
	t := time.Now().Add(time.Minute * 6).Unix()
	deadline := big.NewInt(t)

	addApprove(&tos, &data, ammountIn[0], collateralAdress, forgeRouterAddress)
	addSwap(&tos, &data, ammountIn[0], path, contractAdress, deadline, forgeRouterAddress, forgeRouterContract)
	addLiquidation(&tos, &data, accountToLiquidate, assetID, ammountToLiquidate, forgeAddress, forgeContract)

	multicallParsed, err := abi.JSON(strings.NewReader(multicall.MulticallABI))
	if err != nil {
		fmt.Println("error parsing abi multicall")
		log.Fatal(err)
	}

	params, err := multicallParsed.Pack("swap", tos, data)
	if err != nil {
		fmt.Println("error packing multicall")
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA(privateKeyString)
	if err != nil {
		log.Fatal(err)
	}

	fromAddress := common.HexToAddress(publicKeyString)

	nonce, err := conn.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := conn.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	chainId, err := conn.ChainID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
	if err != nil {
		log.Fatal(err)
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0) // in wei
	auth.GasLimit = uint64(0)  // in units
	auth.GasPrice = gasPrice

	flashLoanContract, err := flashloan.NewFlashloan(contractAdress, conn)
	if err != nil {
		log.Fatal(err)
	}

	tx, err := flashLoanContract.AaveFlashloan(auth, collateralAdress, ammountIn[0], params[4:])
	if err != nil {
		fmt.Println("error sending tx")
		discord.SendErrorMessage(err, discordUrl)
		log.Fatal(err)
	}
	skip.Sub(skip, big.NewInt(1))
	fmt.Printf("tx sent: %s", tx.Hash().Hex())
	discord.SendDiscordMessage(syntethicName, ammountIn[0], tx.Hash().Hex(), discordUrl)
	fmt.Println("\n Time for sending a liquidation: ", time.Since(start))

}

func main() {
	if err := env.Load("config.env"); err != nil {
		panic(err)
	}
	privateKeyString := os.Getenv("PRIVATE_KEY")
	publicKeyString := os.Getenv("PUBLIC_KEY")
	discordUrl := os.Getenv("DISCORD_URL")
	provider := os.Getenv("PROVIDER")
	contractAdress := common.HexToAddress(os.Getenv("CONTRACT_ADDRESS"))

	client, err := ethclient.Dial(provider)
	if err != nil {
		log.Fatal(err)
	}

	liquidationFinderAdress := common.HexToAddress("0x9F0d16dbB77f8Ec5dd2eeb021659f87b56a404Be")
	liquidationFinderContract, err := liquidationFinder.NewLiquidationFinder(liquidationFinderAdress, client)
	if err != nil {
		log.Fatal(err)
	}
	usdcToken := common.HexToAddress("0xFF970A61A04b1cA14834A43f5dE4533eBDDB5CC8")
	xauToken := common.HexToAddress("0xB03608FAfA7908e9C8FCbea2A5a3c66C2D4aBbb1")
	cnyToken := common.HexToAddress("0x7F7Ef19bdb8dFB4d8117D606c2A97efE1BcE4B0b")
	wtiToken := common.HexToAddress("0xd4f3640dD3D7332Fb2cAD672089F97DFe6D812CA")
	xagToken := common.HexToAddress("0xbdA42696270e0985625Ce27f83667f0aB926dAF3")
	eurToken := common.HexToAddress("0x93538304690E0d4881A3cDc47120562da68C87c6")
	jpyToken := common.HexToAddress("0xfB93cd576AeD0be0bAC822A3399285c73D195cA7")
	tslaToken := common.HexToAddress("0xD0B8e2b5716540487D037Abfd5521E0f4d2BBB2C")

	xauSkip := big.NewInt(0)
	cnySkip := big.NewInt(0)
	wtiSkip := big.NewInt(0)
	xagSkip := big.NewInt(0)
	eurSkip := big.NewInt(0)
	jpySkip := big.NewInt(0)
	tslaSkip := big.NewInt(0)

	for true {
		resXau := findXauLiquidation(liquidationFinderContract, xauSkip)
		fmt.Println(resXau)
		if resXau.Collateral.Cmp(big.NewInt(0)) == 1 {
			fmt.Println("Found liquidation for XAU")
			go flashloanAndLiquidate(resXau.Amount, client, usdcToken, xauToken, contractAdress, big.NewInt(0), resXau.Account, "XAU", xauSkip, privateKeyString, publicKeyString, discordUrl)
			xauSkip = xauSkip.Add(xauSkip, big.NewInt(1))
		} else {
			fmt.Println("No liquidation found in XAU")
		}
		resCny := findCnyLiquidation(liquidationFinderContract, cnySkip)
		fmt.Println(resCny)
		if resCny.Collateral.Cmp(big.NewInt(0)) == 1 {
			fmt.Println("Found liquidation for CNY")
			go flashloanAndLiquidate(resCny.Amount, client, usdcToken, cnyToken, contractAdress, big.NewInt(1), resCny.Account, "CNY", cnySkip, privateKeyString, publicKeyString, discordUrl)
			cnySkip = cnySkip.Add(cnySkip, big.NewInt(1))
		} else {
			fmt.Println("No liquidation found in CNY")
		}
		resWti := findWtiLiquidation(liquidationFinderContract, wtiSkip)
		fmt.Println(resWti)
		if resWti.Collateral.Cmp(big.NewInt(0)) == 1 {
			fmt.Println("Found liquidation for WTI")
			go flashloanAndLiquidate(resWti.Amount, client, usdcToken, wtiToken, contractAdress, big.NewInt(2), resWti.Account, "WTI", wtiSkip, privateKeyString, publicKeyString, discordUrl)
			wtiSkip = wtiSkip.Add(wtiSkip, big.NewInt(1))
		} else {
			fmt.Println("No liquidation found in WTI")
		}
		resXag := findXagLiquidation(liquidationFinderContract, xagSkip)
		fmt.Println(resXag)
		if resXag.Collateral.Cmp(big.NewInt(0)) == 1 {
			fmt.Println("Found liquidation for XAG")
			go flashloanAndLiquidate(resXag.Amount, client, usdcToken, xagToken, contractAdress, big.NewInt(3), resXag.Account, "XAG", xagSkip, privateKeyString, publicKeyString, discordUrl)
			xagSkip = xagSkip.Add(xagSkip, big.NewInt(1))
		} else {
			fmt.Println("No liquidation found in XAG")
		}
		resEur := findEurLiquidation(liquidationFinderContract, eurSkip)
		fmt.Println(resEur)
		if resEur.Collateral.Cmp(big.NewInt(0)) == 1 {
			fmt.Println("Found liquidation for EUR")
			go flashloanAndLiquidate(resEur.Amount, client, usdcToken, eurToken, contractAdress, big.NewInt(4), resEur.Account, "EUR", eurSkip, privateKeyString, publicKeyString, discordUrl)
			eurSkip = eurSkip.Add(eurSkip, big.NewInt(1))
		} else {
			fmt.Println("No liquidation found in EUR")
		}
		resJpy := findJpyLiquidation(liquidationFinderContract, jpySkip)
		fmt.Println(resJpy)
		if resJpy.Collateral.Cmp(big.NewInt(0)) == 1 {
			fmt.Println("Found liquidation for JPY")
			go flashloanAndLiquidate(resJpy.Amount, client, usdcToken, jpyToken, contractAdress, big.NewInt(5), resJpy.Account, "JPY", jpySkip, privateKeyString, publicKeyString, discordUrl)
			jpySkip = jpySkip.Add(jpySkip, big.NewInt(1))
		} else {
			fmt.Println("No liquidation found in JPY")
		}
		resTsla := findTslaLiquidation(liquidationFinderContract, tslaSkip)
		fmt.Println(resTsla)
		if resTsla.Collateral.Cmp(big.NewInt(0)) == 1 {
			fmt.Println("Found liquidation for TSLA")
			go flashloanAndLiquidate(resTsla.Amount, client, usdcToken, tslaToken, contractAdress, big.NewInt(6), resTsla.Account, "TSLA", tslaSkip, privateKeyString, publicKeyString, discordUrl)
			tslaSkip = tslaSkip.Add(tslaSkip, big.NewInt(1))
		} else {
			fmt.Println("No liquidation found in TSLA")
		}
		time.Sleep(1 * time.Second)

	}
}
