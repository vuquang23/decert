import { CertificateCollection } from "api/certificate-collections";
import Center from "components/Center";
import { useMetaMask } from "components/MetaMaskProvider";
import { useState } from "react";
import { useLocation } from "react-router-dom";

const NewCertificatePage = () => {
  const location = useLocation();
  const collection = location.state as CertificateCollection;
  const [image, setImage] = useState<File>();
  const metaMask = useMetaMask();
  return (
    <>
      <h1 className="display-4 mb-5">Issue new "{collection.title}"</h1>
      <form>
        <div className="row g-5">
          <div className="col-12 col-md-8">
            <div className="row g-3">
              <Issuer address={metaMask.address} />
              <hr className="border opacity-100" />
              <Receiver />
              <Certificate />
              <div className="d-none d-md-block col-12">
                <button type="submit" className="btn btn-primary">
                  Issue
                </button>
              </div>
            </div>
          </div>
          <ImageInput image={image} setImage={(image) => setImage(image)} />
        </div>
        <div className="row mt-3 mb-5">
          <div className="d-block d-md-none col-12">
            <button type="submit" className="btn btn-primary">
              Issue
            </button>
          </div>
        </div>
      </form>
    </>
  );
};

const Issuer = ({ address }: { address: string }) => {
  return (
    <>
      <div className="col-12">
        <label className="form-label">Issuer Address</label>
        <input
          type="text"
          className="form-control"
          readOnly
          disabled
          value={address}
        />
      </div>
      <div className="col-md-6">
        <label className="form-label">Issuer name</label>
        <input type="text" className="form-control" />
      </div>
      <div className="col-md-6 mb-5">
        <label className="form-label">Issuer's position</label>
        <input type="text" className="form-control" />
      </div>
    </>
  );
};

const Receiver = () => (
  <>
    <div className="col-12">
      <label className="form-label">Receiver Address</label>
      <input type="text" className="form-control" />
    </div>
    <div className="col-md-6">
      <label className="form-label">Receiver name</label>
      <input type="text" className="form-control" />
    </div>
    <div className="col-md-6">
      <label className="form-label">Date of birth</label>
      <input type="date" className="form-control" />
    </div>
  </>
);

const Certificate = () => (
  <>
    <div className="col-md-6">
      <label className="form-label">Expired at</label>
      <input type="date" className="form-control" />
    </div>
    <div className="col-12">
      <label className="form-label">Description</label>
      <textarea className="form-control" rows={5} />
    </div>
  </>
);

const ImageInput = ({
  image,
  setImage,
}: {
  image?: File;
  setImage: (image?: File) => void;
}) => (
  <div
    className="col-12 col-md-4"
    onDragOver={(event) => event.preventDefault()}
    onDrop={(event) => {
      event.preventDefault();
      setImage(event.dataTransfer.files[0]);
    }}
  >
    <Center className="h-100 w-100 border rounded flex-column position-relative">
      {typeof image !== "undefined" ? (
        <>
          <img
            src={URL.createObjectURL(image)}
            alt="Certificate"
            className="rounded mw-100"
          />
          <label
            htmlFor="image"
            className="btn btn-light shadow border rounded-pill position-absolute bottom-0 start-50 translate-middle-x mb-3"
          >
            Change
          </label>
        </>
      ) : (
        <Center className="flex-column my-5">
          <i className="h1 text-secondary bi bi-file-image" />
          <h6 className="text-secondary mb-3">Certificate picture</h6>
          <label htmlFor="image" className="btn btn-sm btn-outline-secondary">
            Upload
          </label>
        </Center>
      )}
      <input
        type="file"
        id="image"
        hidden
        onChange={(event) => {
          event.preventDefault();
          const files = event.target.files;
          if (files !== null) {
            setImage(files[0]);
          }
        }}
      />
    </Center>
  </div>
);

export default NewCertificatePage;
