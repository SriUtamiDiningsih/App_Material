// Toko Material POS - Cashier Interactive Script

let cartItems = [
  { id: 1, name: "Semen Tiga Roda 50 Kg", price: 70000, qty: 2, unit: "Zak", discount: 0 },
  { id: 2, name: "Besi Beton 12mm", price: 85000, qty: 5, unit: "Batang", discount: 0 },
  { id: 3, name: "Pasir Pasang", price: 350000, qty: 1, unit: "m³", discount: 0 }
];

const TAX_RATE = 0.11; // 11% PPN

function formatRupiah(number) {
  return new Intl.NumberFormat("id-ID", {
    style: "currency",
    currency: "IDR",
    minimumFractionDigits: 0,
    maximumFractionDigits: 0
  }).format(number).replace("IDR", "Rp");
}

function renderCart() {
  const tbody = document.getElementById("cartTableBody");
  if (!tbody) return;

  tbody.innerHTML = "";
  let subtotal = 0;

  cartItems.forEach((item, index) => {
    const itemSubtotal = (item.price * item.qty) - (item.discount || 0);
    subtotal += itemSubtotal;

    const tr = document.createElement("tr");
    tr.innerHTML = `
      <td class="text-muted fw-semibold">${index + 1}</td>
      <td>
        <div class="fw-semibold text-dark">${item.name}</div>
      </td>
      <td>${formatRupiah(item.price)}</td>
      <td>
        <div class="input-group input-group-sm" style="width: 105px;">
          <button class="btn btn-outline-secondary px-2" type="button" onclick="updateQty(${index}, -1)">-</button>
          <input type="number" class="form-control text-center fw-bold" value="${item.qty}" min="1" onchange="setQty(${index}, this.value)">
          <button class="btn btn-outline-secondary px-2" type="button" onclick="updateQty(${index}, 1)">+</button>
        </div>
      </td>
      <td><span class="badge bg-light text-dark border">${item.unit}</span></td>
      <td>
        <input type="number" class="form-control form-control-sm text-end" style="width: 80px;" value="${item.discount}" onchange="setDiscount(${index}, this.value)">
      </td>
      <td class="fw-bold text-dark">${formatRupiah(itemSubtotal)}</td>
      <td class="text-end">
        <button class="btn btn-action delete" onclick="removeItem(${index})" title="Hapus Item">
          <i class="bi bi-trash"></i>
        </button>
      </td>
    `;
    tbody.appendChild(tr);
  });

  const totalDiscount = cartItems.reduce((acc, it) => acc + (Number(it.discount) || 0), 0);
  const taxAmount = Math.round(subtotal * TAX_RATE);
  const grandTotal = subtotal + taxAmount;

  document.getElementById("summarySubtotal").innerText = formatRupiah(subtotal);
  document.getElementById("summaryDiskon").innerText = formatRupiah(totalDiscount);
  document.getElementById("summaryPajak").innerText = formatRupiah(taxAmount);
  document.getElementById("summaryTotal").innerText = formatRupiah(grandTotal);

  calculateChange(grandTotal);
}

function updateQty(index, delta) {
  if (cartItems[index].qty + delta > 0) {
    cartItems[index].qty += delta;
    renderCart();
  }
}

function setQty(index, val) {
  const num = parseInt(val);
  if (num > 0) {
    cartItems[index].qty = num;
    renderCart();
  }
}

function setDiscount(index, val) {
  const num = parseFloat(val) || 0;
  cartItems[index].discount = num;
  renderCart();
}

function removeItem(index) {
  cartItems.splice(index, 1);
  renderCart();
}

function calculateChange(currentGrandTotal) {
  const grandTotal = currentGrandTotal !== undefined ? currentGrandTotal : getCurrentGrandTotal();
  const inputBayar = document.getElementById("inputBayar");
  const displayKembali = document.getElementById("displayKembali");

  if (!inputBayar || !displayKembali) return;

  const bayarValue = parseFloat(inputBayar.value) || 0;
  const kembalian = bayarValue - grandTotal;

  if (kembalian >= 0) {
    displayKembali.innerText = formatRupiah(kembalian);
    displayKembali.className = "val text-success";
  } else {
    displayKembali.innerText = "-" + formatRupiah(Math.abs(kembalian));
    displayKembali.className = "val text-danger";
  }
}

function getCurrentGrandTotal() {
  let subtotal = 0;
  cartItems.forEach(item => {
    subtotal += (item.price * item.qty) - (item.discount || 0);
  });
  return subtotal + Math.round(subtotal * TAX_RATE);
}

function addProductToCart(product) {
  const existing = cartItems.find(it => it.id === product.id);
  if (existing) {
    existing.qty += 1;
  } else {
    cartItems.push({
      id: product.id,
      name: product.name,
      price: product.price,
      qty: 1,
      unit: product.unit,
      discount: 0
    });
  }
  renderCart();
}

function quickAddSample() {
  const sampleProducts = [
    { id: 4, name: "Bata Merah Super", price: 950, unit: "Pcs" },
    { id: 5, name: "Cat Dulux Putih 5kg", price: 175000, unit: "Kaleng" },
    { id: 6, name: "Pipa PVC 4 inch Wavin", price: 68000, unit: "Batang" }
  ];
  const randomProd = sampleProducts[Math.floor(Math.random() * sampleProducts.length)];
  addProductToCart(randomProd);
}

function resetCart() {
  if (confirm("Apakah Anda yakin ingin membatalkan transaksi ini?")) {
    cartItems = [];
    renderCart();
  }
}

function processPayment() {
  const grandTotal = getCurrentGrandTotal();
  const inputBayar = document.getElementById("inputBayar");
  const bayarValue = parseFloat(inputBayar ? inputBayar.value : 0) || 0;

  if (cartItems.length === 0) {
    alert("Keranjang masih kosong!");
    return;
  }

  if (bayarValue < grandTotal) {
    alert("Jumlah pembayaran kurang dari total tagihan!");
    return;
  }

  const kembali = bayarValue - grandTotal;
  alert(`Transaksi Berhasil Disimpan!\n\nNo Faktur: TRX-2024-05-21-0001\nTotal: ${formatRupiah(grandTotal)}\nBayar: ${formatRupiah(bayarValue)}\nKembali: ${formatRupiah(kembali)}`);
  
  // Reset cart
  cartItems = [];
  renderCart();
  if (inputBayar) inputBayar.value = "";
}

document.addEventListener("DOMContentLoaded", () => {
  renderCart();
  const inputBayar = document.getElementById("inputBayar");
  if (inputBayar) {
    inputBayar.addEventListener("input", () => calculateChange());
  }
});