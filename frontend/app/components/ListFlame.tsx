import '../styles/ListFlame.scss';
import {useState, useEffect} from "react";

const ListFlame = ({data}) => {

  const [index, setIndex] = useState(0);
  const prevItem = () => setIndex((i) => Math.max(i - 1, 0));
  const nextItem = () => setIndex((i) => Math.min(i + 1, data.length - 1));

  useEffect(() => {
    const handleKeyDown = (event: KeyboardEvent) => {
      if (event.key === "ArrowLeft") prevItem();
      if (event.key === "ArrowRight") nextItem();
    };
    window.addEventListener("keydown", handleKeyDown);
    return () => window.removeEventListener("keydown", handleKeyDown);
  }, []);

  return (
    <div className="listflame">
      <ul>
        {Object.entries(data[index]).map(([key, value]) => (
          <li key={key}>
            <div className="key">{key}</div>
            <div className="value">{value}</div>
          </li>
        ))}
      </ul>

      <div className="navigation">
        <button onClick={prevItem} disabled={index === 0}>前へ</button>

        <div className="id-selector">
          <span>IDの値:</span>
          <select
            value={index}
            onChange={(e) => setIndex(parseInt(e.target.value, 10))}
          >
            {data.map((_, i) => (
              <option key={i} value={i}>{i + 1}</option>
            ))}
          </select>
        </div>

        <button onClick={nextItem} disabled={index === data.length - 1}>次へ</button>
      </div>
    </div>
  );
};

export default ListFlame;
