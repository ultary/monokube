from django.shortcuts import render

# Create your views here.


def index(request):
    context = None
    template_name = 'kluster/index.html'
    return render(request, template_name, context)
