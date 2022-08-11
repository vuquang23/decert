import Swal from "sweetalert2";

const BootstrapSwal = Swal.mixin({
  customClass: {
    // container: "...",
    // popup: "...",
    title: "h3",
    // closeButton: "...",
    // icon: "...",
    // image: "...",
    // htmlContainer: "...",
    input: "form-control w-auto",
    inputLabel: "form-label",
    validationMessage: "invalid-feedback",
    // actions: "...",
    confirmButton: "btn btn-success mx-1",
    denyButton: "btn btn-danger mx-1",
    cancelButton: "btn btn-secondary mx-1",
    loader: "border-0 animation-none",
    // footer: "....",
    // timerProgressBar: "....",
  },
  backdrop: true,
  buttonsStyling: false,
  loaderHtml: '<div class="spinner-border" />',
  closeButtonHtml: '<i class="bi bi-x" />',
});

const BootstrapSwalDanger = Swal.mixin({
  customClass: {
    // container: "...",
    // popup: "...",
    title: "h3",
    // closeButton: "...",
    // icon: "...",
    // image: "...",
    // htmlContainer: "...",
    input: "form-control w-auto",
    inputLabel: "form-label",
    validationMessage: "invalid-feedback",
    // actions: "...",
    confirmButton: "btn btn-danger mx-1",
    denyButton: "btn btn-success mx-1",
    cancelButton: "btn btn-secondary mx-1",
    loader: "border-0 animation-none",
    // footer: "....",
    // timerProgressBar: "....",
  },
  backdrop: true,
  buttonsStyling: false,
  loaderHtml: '<div class="spinner-border" />',
  closeButtonHtml: '<i class="bi bi-x" />',
});

export { BootstrapSwal, BootstrapSwalDanger };
