import { ReactNode } from "react";

const Center = ({
  className,
  children,
}: {
  className: string;
  children: ReactNode;
}) => (
  <div
    className={`d-flex align-items-center justify-content-center ${className}`}
  >
    {children}
  </div>
);

export default Center;
