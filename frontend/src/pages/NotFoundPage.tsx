import image from "assets/404.jpg";
import { ErrorPage } from "./ErrorPage";

const NotFoundPage = () => (
  <ErrorPage
    image={image}
    attrUrl="https://www.freepik.com/vectors/server-error"
    attrContent="Server error vector created by storyset - www.freepik.com"
  />
);

export default NotFoundPage;
