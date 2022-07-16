import NavBar from "../components/NavBar";

const HomePage = () => (
  <div className="vh-100">
    <NavBar transparent />
    <MainContent />
    <ImageAttribution />
  </div>
);

const MainContent = () => (
  <div className="h-100 w-100 fixed-top z-index-1 homepage-bg">
    <div className="container position-absolute top-50 start-50 translate-middle">
      <h1 className="display-2 text-light text-center mb-4">Decert Verifier</h1>
      <div className="input-group input-group-lg">
        <input
          type="text"
          className="form-control"
          placeholder="Certificate URL"
        />
        <button className="btn btn-success" type="button">
          <span className="mx-3 fw-bold">Verify</span>
        </button>
      </div>
    </div>
  </div>
);

const ImageAttribution = () => {
  const authorUrl =
    "https://unsplash.com/@hjrc33?utm_source=unsplash&utm_medium=referral&utm_content=creditCopyText";
  const imageUrl =
    "https://unsplash.com/s/photos/business?utm_source=unsplash&utm_medium=referral&utm_content=creditCopyText";
  const iconUrl = "https://www.flaticon.com/free-icons/global";

  return (
    <p className="position-fixed bottom-0 end-0 me-2 z-index-2 text-light small opacity-50">
      Photo by{" "}
      <a className="link-light" href={authorUrl}>
        Héctor J. Rivas
      </a>{" "}
      on{" "}
      <a className="link-light" href={imageUrl}>
        Unsplash
      </a>{" "}
      &bull;{" "}
      <a className="link-light" href={iconUrl}>
        Global icons created by Freepik - Flaticon
      </a>
    </p>
  );
};

export default HomePage;
