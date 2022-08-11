import { getShortAddress } from "components/MetaMaskProvider";

const Address = ({
  address,
  customTextClassName,
}: {
  address: string;
  customTextClassName?: string;
}) => (
  <>
    <code className={customTextClassName}>{getShortAddress(address)}</code>{" "}
    <button
      type="button"
      className="btn btn-light"
      onClick={() => navigator.clipboard.writeText(address)}
    >
      <i className="bi bi-clipboard align-baseline" />
    </button>
  </>
);

export default Address;
