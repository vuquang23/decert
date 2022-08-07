import image from "assets/404.jpg";
import { ErrorPage } from "./ErrorPage";

const NotFoundPage = () => (
  <ErrorPage
    image={image}
    attrUrl="https://www.freepik.com/vectors/server-error"
    attrContent="Server error vector created by storyset - www.freepik.com"
  />
);

class NotFoundError extends Error {}

export default NotFoundPage;
export { NotFoundError };
