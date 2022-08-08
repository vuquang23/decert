import { Certificate, issue } from "api/certificate";
import { CertificateCollection } from "api/certificate-collections";
import { BootstrapSwal } from "components/BootstrapSwal";
import Center from "components/Center";
import { useMetaMask } from "components/MetaMaskProvider";
import {
  dateFromYYYYMMDD,
  formValidationClassName,
  toDDMMYYYYstring,
  userRejectTransaction,
} from "helper";
import { onPromiseRejected } from "pages/ErrorPage";
import { useState } from "react";
import { FormProvider, useForm, useFormContext } from "react-hook-form";
import { useLocation, useNavigate } from "react-router-dom";
import Swal from "sweetalert2";

const NewCertificatePage = () => {
  const location = useLocation();
  const collection = location.state as CertificateCollection;
  const [image, setImage] = useState<File>();
  const metaMask = useMetaMask();
  const form = useForm<Certificate>({
    defaultValues: {
      id: 0,
      certNftId: 0,
      collectionId: collection.id,
      certTitle: collection.collectionName,
      certImage: "",
      issuedAt: Date.now(),
    },
  });
  const navigate = useNavigate();

  const submitHandler = (cert: Certificate) => {
    BootstrapSwal.fire({
      icon: "question",
      title: "Do you want to issue this certificate?",
      showConfirmButton: true,
      showCancelButton: true,
      showLoaderOnConfirm: true,
      preConfirm: () => issue(metaMask, collection, cert),
      allowOutsideClick: () => !Swal.isLoading(),
    })
      .then(() =>
        BootstrapSwal.fire({
          icon: "success",
          title: "Transaction broadcast!",
        })
      )
      .then(() => navigate("/collections"))
      .catch((reason) => {
        Swal.close();
        if (userRejectTransaction(reason)) {
          BootstrapSwal.fire({
            icon: "error",
            title: "You rejected the transaction.",
          });
        } else {
          onPromiseRejected(reason, navigate);
        }
      });
  };

  return (
    <>
      <h1 className="display-4 mb-5">
        Issue new "{collection.collectionName}"
      </h1>
      <FormProvider {...form}>
        <form onSubmit={form.handleSubmit(submitHandler)}>
          <div className="row g-5">
            <div className="col-12 col-md-8">
              <div className="row g-3">
                <IssuerInfo address={metaMask.address} />
                <hr className="border opacity-100" />
                <ReceiverInfo />
                <CertificateInfo />
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
      </FormProvider>
    </>
  );
};

const IssuerInfo = ({ address }: { address: string }) => {
  const {
    register,
    formState: { errors },
  } = useFormContext<Certificate>();
  return (
    <>
      <div className="col-12">
        <label className="form-label">Issuer Address</label>
        <input
          type="text"
          className="form-control"
          readOnly
          disabled
          {...register("issuer.wallet", {
            required: true,
            value: address,
          })}
        />
      </div>
      <div className="col-md-6">
        <label className="form-label">Issuer name</label>
        <input
          type="text"
          className={`form-control ${formValidationClassName(
            errors.issuer?.name
          )}`}
          {...register("issuer.name", { required: true })}
        />
        <div className="invalid-feedback">Issuer name is required</div>
      </div>
      <div className="col-md-6 mb-5">
        <label className="form-label">Issuer's position</label>
        <input
          type="text"
          className={`form-control ${formValidationClassName(
            errors.issuer?.position
          )}`}
          {...register("issuer.position", { required: true })}
        />
        <div className="invalid-feedback">Issuer's position is required</div>
      </div>
    </>
  );
};

const ReceiverInfo = () => {
  const {
    register,
    formState: { errors },
  } = useFormContext<Certificate>();
  return (
    <>
      <div className="col-12">
        <label className="form-label">Receiver Address</label>
        <input
          type="text"
          className={`form-control ${formValidationClassName(
            errors.receiver?.wallet
          )}`}
          {...register("receiver.wallet", {
            required: true,
            pattern: /^0x[a-fA-F0-9]{40}$/g,
          })}
        />
        <div className="invalid-feedback">Invalid address</div>
      </div>
      <div className="col-md-6">
        <label className="form-label">Receiver name</label>
        <input
          type="text"
          className={`form-control ${formValidationClassName(
            errors.receiver?.name
          )}`}
          {...register("receiver.name", { required: true })}
        />
        <div className="invalid-feedback">Receiver name is required</div>
      </div>
      <div className="col-md-6">
        <label className="form-label">Date of birth</label>
        <input
          type="date"
          className={`form-control ${formValidationClassName(
            errors.receiver?.dateOfBirth
          )}`}
          max={new Date(Date.now()).toISOString().split("T")[0]}
          {...register("receiver.dateOfBirth", {
            validate: (date) => date.length > 0,
            setValueAs: (value) => {
              const date = value as string;
              return date.length !== 0
                ? toDDMMYYYYstring(dateFromYYYYMMDD(date, "-"))
                : "";
            },
          })}
        />
        <div className="invalid-feedback">
          Receiver's date of birth is required
        </div>
      </div>
    </>
  );
};

const CertificateInfo = () => {
  const {
    register,
    formState: { errors },
  } = useFormContext<Certificate>();
  return (
    <>
      <div className="col-md-6">
        <label className="form-label">Expired at</label>
        <input
          type="date"
          className="form-control"
          min={new Date(Date.now()).toISOString().split("T")[0]}
          {...register("expiredAt", {
            setValueAs: (value) => {
              const date = value as string;
              return date.length !== 0
                ? dateFromYYYYMMDD(date, "-").getTime()
                : "null";
            },
          })}
        />
      </div>
      <div className="col-12">
        <label className="form-label">Description</label>
        <textarea
          className={`form-control ${formValidationClassName(
            errors.description
          )}`}
          rows={5}
          {...register("description", { required: true })}
        />
        <div className="invalid-feedback">Description is required</div>
      </div>
    </>
  );
};

const ImageInput = ({
  image,
  setImage,
}: {
  image?: File;
  setImage: (image?: File) => void;
}) => {
  const {
    register,
    setValue,
    formState: { errors },
  } = useFormContext<Certificate>();

  const textColor = errors.imgFiles !== undefined ? "danger" : "secondary";
  return (
    <div
      className="col-12 col-md-4"
      onDragOver={(event) => event.preventDefault()}
      onDrop={(event) => {
        event.preventDefault();
        setValue("imgFiles", event.dataTransfer.files, {
          shouldValidate: true,
        });
        setImage(event.dataTransfer.files[0]);
      }}
    >
      <Center
        className={`h-100 w-100 border rounded flex-column position-relative ${
          errors.imgFiles !== undefined ? "border-danger" : ""
        }`}
      >
        {image !== undefined ? (
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
            <i className={`h1 text-${textColor} bi bi-file-image`} />
            <h6 className={`text-${textColor} mb-3`}>Certificate picture</h6>
            <label
              htmlFor="image"
              className={`btn btn-sm btn-outline-${textColor}`}
            >
              Upload
            </label>
          </Center>
        )}
        <input
          type="file"
          id="image"
          hidden
          {...register("imgFiles", {
            validate: (value) => value !== undefined && value.length > 0,
            onChange: (event: React.ChangeEvent<HTMLInputElement>) => {
              event.preventDefault();
              const files = event.target.files;
              if (files !== null) {
                setImage(files[0]);
              }
            },
          })}
        />
      </Center>
    </div>
  );
};

export default NewCertificatePage;
