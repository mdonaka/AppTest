import '../styles/QuizFlame.scss';
import {useState, useEffect} from "react";

const QuizFlame = ({data, editableFields}) => {
  const [index, setIndex] = useState(0);
  const [formData, setFormData] = useState({});

  useEffect(() => {
    if (data.length > 0) {
      const initialData = {...data[index]};
      editableFields.forEach(field => {
        if (initialData[field] !== undefined) {
          initialData[field] = "";
        }
      });
      setFormData(initialData);
    }
  }, [data, index, editableFields]);

  const handleChange = (e) => {
    const {name, value} = e.target;
    setFormData({...formData, [name]: value});
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    const queryParams = new URLSearchParams(formData).toString();
    const url = `http://localhost:8080/check?${queryParams}`;
    try {
      const response = await fetch(url);
      const result = await response.json();
      console.log(result);
    } catch (error) {
      console.error('Error:', error);
    }
  };

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
    <div className="quizflame">
      <form onSubmit={handleSubmit}>
        <ul>
          {Object.entries(formData).map(([key, value]) => (
            <li key={key}>
              <div className="key">{key}</div>
              <div className="value">
                {editableFields.includes(key) ? (
                  <input
                    type="text"
                    id={key}
                    name={key}
                    value={value}
                    onChange={handleChange}
                  />
                ) : (
                  <span>{value}</span>
                )}
              </div>
            </li>
          ))}
        </ul>
      </form>

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
        <button type="submit" className="submit-button">Submit</button>
      </div>
    </div>
  );
};

export default QuizFlame;
