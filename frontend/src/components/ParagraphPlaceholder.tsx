import { arrayFromSize } from "helper";

const ParagraphPlaceholder = () => (
  <p>
    {arrayFromSize(5, (index) => (
      <span key={index} className={`placeholder col-${getRandomCol()} me-2`} />
    ))}
  </p>
);

const getRandomCol = () => Math.floor(Math.random() * 12) + 1;

export default ParagraphPlaceholder;
export { getRandomCol };
