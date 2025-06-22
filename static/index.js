(async () => {
    const platform = navigator.platform;
    const language = navigator.language;

    // Get geolocation as a Promise
    const geolocationPromise = new Promise((resolve, reject) => {
	navigator.geolocation.getCurrentPosition(
	    pos => resolve({ lat: pos.coords.latitude, lon: pos.coords.longitude }),
	    err => reject(err)
	);
    });

    // Fetch IP address as a promise
    const ipPromise = fetch('https://api.ipify.org?format=json')
	  .then(res => res.json())
	  .then(data => data.ip);

    try {
	const [{ lat, lon }, IP] = await Promise.all([geolocationPromise, ipPromise]);
	const payload = { platform, language, IP, lat, lon };
	console.log('Sending:', payload);

	await fetch('http://localhost:6969/send', {
	    method: 'POST',
	    headers: { 'Content-Type': 'application/json' },
	    body: JSON.stringify(payload)
	});

	console.log('Data successfully sent to backend.');
    } catch (error) {
	console.error('Error:', error);
    }


})();
