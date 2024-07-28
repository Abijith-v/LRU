import './App.css';
import Set from './components/Set.jsx'
import Get from './components/Get.jsx'
import Delete from './components/Delete.jsx'

function App() {
  return (
    <div className="App">
      <div className='container'>
      <Set/>
      <Get/>
      <Delete/>
      </div>
    </div>
  );
}

export default App;
