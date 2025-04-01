import {useState, useEffect} from 'react';
import QuizFlame from '../components/QuizFlame';
import '../styles/Quiz.scss';

export default function Quiz() {
  const [data, setData] = useState([]);
  const [editableFields, setEditableFields] = useState<string[]>([]);

  useEffect(() => {
    async function fetchData() {
      const response = await fetch('http://localhost:8080/data');
      const result = await response.json();
      setData(result);
    }
    fetchData();
  }, []);

  const handleFieldChange = (field) => {
    setEditableFields(prevFields =>
      prevFields.includes(field)
        ? prevFields.filter(f => f !== field)
        : [...prevFields, field]
    );
  };

  return (
    <div className="quiz-page">
      {data.length === 0 ? (
        <h1>Loading...</h1>
      ) : (
        <>
          <div className="field-selector-container">
            <h2>Select Quiz Fields</h2>
            <div className="field-selector">
              {Object.keys(data[0] || {}).map((key) => (
                (key !== 'id' && key !== 'name') && (
                  <label key={key} className="custom-checkbox">
                    <input
                      type="checkbox"
                      checked={editableFields.includes(key)}
                      onChange={() => handleFieldChange(key)}
                    />
                    <span className="checkmark"></span>
                    {key}
                  </label>
                )
              ))}
            </div>
          </div>
          <QuizFlame data={data} editableFields={editableFields} />
        </>
      )}
    </div>
  );
}
