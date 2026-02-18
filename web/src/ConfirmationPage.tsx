import { API_URL } from "./App";

const ConfirmationPage = () => {
  const token = ""
  const handleConfirm = async () => {
    const response = await fetch(`${API_URL}/users/activate/${token}`, {
      method: 'PUT',
    });
    if (response.ok) {
      // redirect to the "/" page
    } else {
      // handle error
      alert('Failed to confirm. Please try again.');
    }
  };

  return (
    <div>
      <h1>Confirmation</h1>
      <button onClick={handleConfirm}>Click to confirm</button>
    </div>
  );
};

export default ConfirmationPage;
