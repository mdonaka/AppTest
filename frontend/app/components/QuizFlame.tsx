import '../styles/QuizFlame.scss';
import {useState, useEffect} from "react";

const check = async (queryParams) => {
  const url = `http://localhost:8080/check?${queryParams}`;
  try {
    const response = await fetch(url, {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json'
      }
    });
    return await response.json();
  } catch (error) {
    console.error('Error:', error);
    throw error;
  }
};

const QuizFlame = ({data, editableFields}) => {
  const [index, setIndex] = useState(0);
  const [formData, setFormData] = useState({});
  const [resultMessage, setResultMessage] = useState("　");

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

    if (name === 'name') {
      const newIndex = data.findIndex(item => item.name === value);
      if (newIndex !== -1) {
        setIndex(newIndex);
      }
    }
  };

  const handleSubmit = async (e) => {
    e.preventDefault();

    const hasEmptyField = Object.values(formData).some(value => value === "");
    if (hasEmptyField) {
      setResultMessage("すべてのフィールドに入力してください");
      return;
    }

    const queryParams = new URLSearchParams(formData).toString();
    try {
      const result = await check(queryParams);
      if (result.match) {
        setResultMessage("正解！");
      } else {
        setResultMessage("不正解...");
      }
    } catch (error) {
      setResultMessage("エラーが発生しました");
    }
  };

  const prevItem = () => {
    setIndex((i) => Math.max(i - 1, 0));
    setResultMessage("　");
  };

  const nextItem = () => {
    setIndex((i) => Math.min(i + 1, data.length - 1));
    setResultMessage("　");
  };

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
            key !== 'id' && (
              <li key={key}>
                <div className="key">{key}</div>
                <div className="value">
                  {editableFields.includes(key) || key === 'name' ? (
                    <select id={key} name={key} value={value} onChange={handleChange}>
                      {Array.from(new Set(data.map(item => item[key]))).sort().map((optionValue, idx) => (
                        <option key={idx} value={optionValue}>
                          {optionValue}
                        </option>
                      ))}
                    </select>
                  ) : (
                    <span>{value}</span>
                  )}
                </div>
              </li>
            )))}
        </ul>
      </form>

      <div className="navigation">
        <button onClick={prevItem} disabled={index === 0}>前へ</button>
        <button onClick={nextItem} disabled={index === data.length - 1}>次へ</button>
        <button type="submit" className="submit-button" onClick={handleSubmit}>Submit</button>
      </div>

      <div className="result-message">{resultMessage}</div>
    </div>
  );
};

export default QuizFlame;
