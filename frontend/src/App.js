import logo from "./logo.svg";
import "./App.scss";
import { Mainnet, DAppProvider, Sepolia, Goerli } from "@usedapp/core";
import BaseLayout from "./Components/BaseLayout/BaseLayout";
import { BrowserRouter } from "react-router-dom";
import { getDefaultProvider } from "ethers";

export const config = {
  readOnlyChainId: Mainnet.chainId,
  readOnlyUrls: {
    [Mainnet.chainId]: getDefaultProvider("mainnet"),
    [Sepolia.chainId]: getDefaultProvider("sepolia"),
    [Goerli.chainId]: getDefaultProvider("goerli"),
  },
};
function App() {
  return (
    <DAppProvider config={config}>
      <BrowserRouter>
        {" "}
        <BaseLayout />
      </BrowserRouter>
    </DAppProvider>
  );
}

export default App;
