from django import forms


class WordCountForm(forms.Form):
    """
    A form for collecting text input and performing word count validation.
    """

    text = forms.CharField(
        widget=forms.Textarea(
            attrs={
                "placeholder": "Enter your text here",
                "rows": 4,
                "cols": 50,
            }
        ),
        required=True,
        error_messages={"required": "Text field should not be empty."},
    )

    # This function was not required just wanted to show the probability of cleaning
    def clean_text(self):
        """
        Cleans and validates the 'text' field value by stripping leading and trailing whitespace.
        Returns the cleaned text value. (in case we want to have delimeters).
        """
        text = self.cleaned_data.get("text")
        if text:
            text = text.strip()  # Stripping leading and trailing whitespace
        return text
