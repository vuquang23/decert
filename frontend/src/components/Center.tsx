const Center = ({
  className,
  children,
}: {
  className: string;
  children: JSX.Element;
}) => (
  <div
    className={`d-flex align-items-center justify-content-center ${className}`}
  >
    {children}
  </div>
);

export default Center;
