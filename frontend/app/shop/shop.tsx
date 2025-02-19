import axios from "axios";
import { useEffect, useState } from "react";

export function Shop() {
  const [catalogue, setCatalogue] = useState([]);
  const [cart, setCart] = useState([]);
  const [total, setTotal] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  useEffect(() => {
    axios.get(`${import.meta.env.VITE_HTTP_URL}/products`)
      .then(response => {
        setCatalogue(response.data.products);
        setLoading(false);
      })
      .catch(error => {
        setError(error);
        setLoading(false);
      });
  }, []);

  useEffect(() => {
    const newTotal = cart.reduce((total, item) => total + item.price * item.quantity, 0);
    setTotal(newTotal);
  }, [cart]);

  const addToCart = (item) => {
    setCart((prevCart) => {
      const existingItem = prevCart.find(cartItem => cartItem.sku === item.sku);

      if (existingItem) {
        return prevCart.map(cartItem =>
          cartItem.sku === item.sku
            ? { ...cartItem, quantity: cartItem.quantity + 1 }
            : cartItem
        );
      } else {
        return [...prevCart, { ...item, quantity: 1 }];
      }
    });
  };

  const removeFromCart = (item) => {
    setCart((prevCart) => {
      const existingItem = prevCart.find(cartItem => cartItem.sku === item.sku);

      if (existingItem && existingItem.quantity === 1) {
        return prevCart.filter(cartItem => cartItem.sku !== item.sku);
      } else {
        return prevCart.map(cartItem =>
          cartItem.sku === item.sku
            ? { ...cartItem, quantity: cartItem.quantity - 1 }
            : cartItem
        );
      }
    });
  };

  const sendPurchaseRequest = () => {
    axios.put(`${import.meta.env.VITE_HTTP_URL}/purchase`, {items: cart})
      .then(response => {
        console.log(response.data)
        alert(`You paid $${response.data.total_price}!`)
      })
      .catch(error => {
        setError(error);
      });
  };

  if (loading) {
    return <div>Loading...</div>;
  }

  if (error) {
    return <div>Error: {error.message}</div>;
  }

  return (
    <main className="flex flex-col items-center justify-center pt-16 pb-4">
      <header className="flex items-center justify-center w-full">
        <div className="p-4">
          <h1 className="text-4xl">Flip Tech Test</h1>
        </div>
      </header>
      <div className="flex items-center justify-center p-10 min-h-0 space-x-50">
        <div className="flex flex-col items-center justify-center min-h-0">
          <h1 className="mb-5">Catalogue</h1>
          <table className="table-auto border-collapse border border-gray-200">
            <thead>
              <tr>
                <th className="border border-gray-300 px-4 py-2">SKU</th>
                <th className="border border-gray-300 px-4 py-2">Name</th>
                <th className="border border-gray-300 px-4 py-2">Price</th>
                <th className="border border-gray-300 px-4 py-2">Inventory Qty</th>
              </tr>
            </thead>
            <tbody>
              {catalogue.map((item, index) => (
                <tr key={index}>
                <td className="border border-gray-300 px-4 py-2">{item.sku}</td>
                <td className="border border-gray-300 px-4 py-2">{item.name}</td>
                <td className="border border-gray-300 px-4 py-2">${item.price}</td>
                <td className="border border-gray-300 px-4 py-2">{item.inventory}</td>
                <td className="border border-gray-300 px-4 py-2">
                  <button className="px-2 py-1 bg-blue-500 text-white rounded active:bg-blue-700" onClick={() => addToCart(item)}>+</button>
                  <button className="px-2 py-1 bg-red-500 text-white rounded ml-2 active:bg-red-700" onClick={() => removeFromCart(item)}>-</button>
                </td>
              </tr>
              ))}
            </tbody>
          </table>
        </div>
        <div className="flex flex-col items-center justify-center min-h-0">
          <h1 className="mb-5">Cart</h1>
          <table className="table-auto border-collapse border border-gray-200">
            <thead>
              <tr>
                <th className="border border-gray-300 px-4 py-2">SKU</th>
                <th className="border border-gray-300 px-4 py-2">Name</th>
                <th className="border border-gray-300 px-4 py-2">Price</th>
                <th className="border border-gray-300 px-4 py-2">Quantity</th>
              </tr>
            </thead>
            <tbody>
              {cart.map((item, index) => (
                <tr key={index}>
                <td className="border border-gray-300 px-4 py-2">{item.sku}</td>
                <td className="border border-gray-300 px-4 py-2">{item.name}</td>
                <td className="border border-gray-300 px-4 py-2">${item.price}</td>
                <td className="border border-gray-300 px-4 py-2">{item.quantity}</td>
              </tr>
              ))}
            </tbody>
          </table>
          <p className="p-5">Total: ${total}</p>
          <button onClick={sendPurchaseRequest} className="px-2 py-1 bg-green-400 text-white rounded active:bg-green-600">Purchase</button>
        </div>
      </div>
    </main>
  )
}
