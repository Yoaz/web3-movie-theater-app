import { config } from "../../App";
import { useEtherBalance, useEthers } from "@usedapp/core";
import "./ConnectButton.scss";

const ConnectButton = () => {
  // Blockchain
  const { activateBrowserWallet, account, chainId, deactivate } = useEthers();
  const userETHBalance = useEtherBalance(account);

  // Truncate wallet address
  const getEllipsisTxt = (str, n = 6) => {
    if (str) {
      return `${str.slice(0, n)}...${str.slice(str.length - n)}`;
    }
    return "";
  };

  // Truncate eth amount
  const truncateDecimals = (number, digits) => {
    var multiplier = Math.pow(10, digits),
      adjustedNum = number * multiplier,
      truncatedNum = Math[adjustedNum < 0 ? "ceil" : "floor"](adjustedNum);

    return truncatedNum / multiplier;
  };

  // In case there is an account but not a supported chain
  if (chainId && !config.readOnlyUrls[chainId]) {
    return <p>Change to Mainnet or Sepolia testnet.</p>;
  }

  return (
    <div className="walletBtn">
      {account ? (
        <li>
          <button onClick={() => deactivate()}>Disconnect</button>
          <div className="connected-info">
            Connected to {getEllipsisTxt(account)} ETH:{" "}
            {userETHBalance
              ? truncateDecimals(userETHBalance / 10e17, 3)
              : "..."}
          </div>
        </li>
      ) : (
        <li>
          <button className="walletBtn" onClick={() => activateBrowserWallet()}>
            Connect
          </button>
        </li>
      )}
    </div>
  );
};

export default ConnectButton;
