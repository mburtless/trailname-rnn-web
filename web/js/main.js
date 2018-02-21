var form = $('#trailname');
var resultStack = [];
$(form).submit(function(event){
	event.preventDefault();
	var formData = $(form).serialize();
	// Submit the form using AJAX.
	$.ajax({
		type: 'POST',
		url: $(form).attr('action'),
		data: formData
	}).done(function(responseData) {
		// you can see the result from the console
		// tab of the developer tools
		console.log(responseData);
		var namesContent = ""
		var names = responseData.result
		for(var i = 0; i < names.length; i++){
		namesContent += "<br>"+names[i];
		}
		//$("#trailNames").html(responseData.result);
		$("#trailNames").empty().append(namesContent);

	}).fail(function(data) {
		if (data.responseText !== '') {
			$("#trailNames").empty().append(data.responseText);
		} else {
			$("#trailNames").empty().append("Oops! We couldn't grab your trail name at this time, please try again later!");
		}
	});
})
