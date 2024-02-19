# from django.shortcuts import render
# from django.contrib.auth.decorators import login_required

# @login_required(login_url='login')
# def dashboard(request):
#     # Examine request
#     print(request.method)  # HTTP method used for the request


from django.shortcuts import render
from django.contrib.auth.decorators import login_required

@login_required(login_url='login')
def dashboard(request):
    #print(request.user)  # User object representing the currently authenticated user

    sourceNode = type = targetNode = None

    if (request.GET['sourceNode'] == request.GET['targetNode']):
        return render(request, 'hpanel/user_panel.html', {'error_message': "Source and Target nodes cannot be the same."})

    if (request.GET['sourceNode'] == 'Node 1'):
        sourceNode = 'team1-NE-1'
    elif (request.GET['sourceNode'] == 'Node 2'):
        sourceNode = 'team1-NE-2'
    elif (request.GET['sourceNode'] == 'Node 3'):
            sourceNode = 'team1-NE-3'
                
    if (request.GET['targetNode'] == 'Node 1'):
        targetNode = 'team1-NE-1'
    elif (request.GET['targetNode'] == 'Node 2'):
        targetNode = 'team1-NE-2'
    elif (request.GET['targetNode'] == 'Node 3'):
        targetNode = 'team1-NE-3'


    if 'connName' in request.GET:
        # Parameter 'connName' exists, set type to 'CreateInfrastructure'
    
        type = 'CreateInfrastructure'
    elif 'serviceRate' in request.GET:
        # Parameter 'serviceRate' exists, set type to 'CreateService'
        type = 'CreateService'
    elif 'deleteType' in request.GET:
        # Parameter 'deleteType' exists, set type to 'DeleteInfrastructure'
        type = 'DeleteInfrastructure'

    # Use the 'type' variable as needed in your code

    # For debugging, you can print the type
    print("Type:", type)

    if type == 'CreateInfrastructure':
        connectionName = request.GET['connName']
        # Print for debugging
        print(f"CreateInfrastructure Form - Source Node: {sourceNode}, Target Node: {targetNode}, Connection Name: {connectionName}")

    elif type == 'CreateService':
            serviceRate = request.GET['serviceRate']
            numberOfService = request.GET['numberOfService']
            # Print for debugging
            print(f"CreateService Form - Source Node: {sourceNode}, Target Node: {targetNode}, Service Rate: {serviceRate}, Number of Service: {numberOfService}")

    elif type == 'DeleteInfrastructure':
            deleteType = request.GET.getlist('deleteType')  # In case multiple checkboxes are selected
            # Print for debugging
            print(f"DeleteInfrastructure Form - Delete Type: {deleteType}, Source Node: {sourceNode}, Target Node: {targetNode}")

    return render(request, 'hpanel/user_panel.html')