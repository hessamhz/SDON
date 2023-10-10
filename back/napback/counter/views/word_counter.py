from counter.forms import WordCountForm
from django.shortcuts import render


def word_count_view(request):
    """
    View function for word count functionality.
    Renders the form and handles form submission to calculate word count.
    """

    # Checks if form is submitted
    if request.method == "POST":
        form = WordCountForm(request.POST)
        if form.is_valid():
            # Get the submitted text from the form's cleaned data
            text = form.cleaned_data.get("text")
            # Count the number of words in the text
            word_count = len(text.split())
            # Render the counter.html template with the word count
            return render(request, "counter.html", {"word_count": word_count})
        else:
            # Render the counter.html template with the invalid form
            return render(request, "counter.html", {"form": form})
    else:
        # Create a new form instance for GET requests
        form = WordCountForm()

    # Render the counter.html template with the form assuming that it is a GET
    return render(request, "counter.html", {"form": form})
