import { ReactComponent as Image } from "assets/empty.svg";

const Empty = () => (
  <>
    <div className="row justify-content-center">
      <div className="col-6">
        <Image />
        <h6 className="display-6 text-center">There is nothing here.</h6>
      </div>
      <a
        className="link-dark small opacity-50 text-center"
        href="https://storyset.com/web"
      >
        Web illustrations by Storyset
      </a>
    </div>
  </>
);

export default Empty;
