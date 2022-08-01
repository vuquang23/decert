import image from "assets/404.jpg";
import Center from "components/Center";
import { useNavigate } from "react-router-dom";

const NotFoundPage = () => {
  const navigate = useNavigate();
  return (
    <>
      <Center className="vh-100">
        <div className="h-50 d-flex flex-column align-items-center">
          <img src={image} alt="Not found" className="h-75 d-block" />
          <button className="btn btn-primary" onClick={() => navigate("/")}>
            Go to homepage
          </button>
        </div>
      </Center>
      <a
        className="link-dark position-fixed bottom-0 end-0 m-2 small opacity-50"
        href="https://www.freepik.com/vectors/server-error"
      >
        Server error vector created by storyset - www.freepik.com
      </a>
    </>
  );
};

export default NotFoundPage;
