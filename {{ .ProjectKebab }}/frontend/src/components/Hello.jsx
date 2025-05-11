import { useState } from 'react';

function Hello() {
  const [count, setCount] = useState(0);

  return (
    <div className="hello-component">
      <h2>Welcome to {{ .Project }}</h2>
      <p>
        A minimal Go backend with React frontend starter (powered by Bun)
      </p>
      <button onClick={() => setCount((count) => count + 1)}>
        Count is {count}
      </button>
    </div>
  );
}

export default Hello;
