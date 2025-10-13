// Small interactive behaviors: mobile nav and booking form
document.getElementById('year').textContent = new Date().getFullYear();

const navToggle = document.getElementById('navToggle');
const nav = document.getElementById('nav');
navToggle.addEventListener('click', ()=>{
  const expanded = nav.style.display === 'flex';
  nav.style.display = expanded ? '' : 'flex';
});

const form = document.getElementById('bookingForm');
form.addEventListener('submit', function(e){
  e.preventDefault();
  const formData = new FormData(form);
  const pickup = formData.get('pickup').trim();
  const dropoff = formData.get('dropoff').trim();
  const dates = formData.get('dates').trim();
  if(!pickup || !dropoff || !dates){
    alert('Please fill pickup, drop-off, and dates.');
    return;
  }
  // Simulate availability check
  const cartype = formData.get('cartype');
  alert(`Checking availability for ${cartype} from ${pickup} — ${dropoff} (${dates}).`);
  form.reset();
});
